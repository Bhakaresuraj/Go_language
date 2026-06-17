package worker
import (
	"fmt"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/database"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/model"
	"sync"
	"time"
)

func StartBackgroundWorker(db *database.Store) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		fmt.Println("Running Health Check...")
		services, err := db.GetAndRunAllServices()
		if err != nil {
			fmt.Println("Error in fetching services from database in worker ....")
			return
		}
		// fmt.Println("Run all Called")
		var wg sync.WaitGroup
		for _, ser := range services {
			wg.Add(1)
			go func(ser model.Service) {
				defer wg.Done()
				avg_time := time.Now()
				ser.Healthy, ser.StatusCode = ser.CheckHealth()
				responseTime := time.Since(avg_time)
				ser.Checked_at = time.Now()
				ser.Response_time = responseTime.Milliseconds()
				err = db.UpdateServiceStatus(ser)
				if err != nil {
					fmt.Println("Update Error :", err)
				}
			}(ser)

		}
		wg.Wait()
		<-ticker.C
	}
}
