import './cards.css'

import GetAllServices from '../../utils/api'
import { useEffect, useState } from 'react';
export default function Card() {
    const [services, setServices] = useState({
        total: 0,
        Healthy: 0,
        unHealthy: 0,
        Avg_response_time: 0
    });
    useEffect(() => {
        function requiredData(data) {
            const totalResponseTime = data.reduce((acc, el) => {
                return acc + el.Response_time;
            }, 0);
            const avgResponseTime =
                data.length > 0
                    ? totalResponseTime / data.length
                    : 0;
            return {
                total: data.length,
                Healthy: data.filter((el) => el.Healthy).length,
                unHealthy: data.filter((el) => !el.Healthy).length,
                Avg_response_time: avgResponseTime.toFixed(0)
            }
        }

        async function fetchServices() {
            try {
                let data = await GetAllServices();
                data = requiredData(data);
                setServices(data);
            } catch (err) {
                console.log(err);
            }
        }
        fetchServices();
        const interval = setInterval(() => {

            fetchServices();

        },10000);
        return () => clearInterval(interval);
    },[]);
    let cardData = [
        {
            icon: "bi bi-link-45deg",
            number: services.total,
            title: "TOTAL SERVICES",
            subtitle: "Monitored endpoints"
        },
        {
            icon: "bi bi-check-square-fill text-success",
            number: services.Healthy,
            title: "HEALTHY SERVICES",
            subtitle: "Operating normally"
        },
        {
            icon: "bi bi-broadcast-pin",
            number: services.unHealthy,
            title: "UNHEALTHY SERVICES",
            subtitle: "Require attention"
        },
        {
            icon: "bi bi-lightning-fill texr",
            number: `${services.Avg_response_time}ms`,
            title: "AVG RESPONSE TIME",
            subtitle: "Last 10 Sec"
        }
    ];


    return (
        <div className="mt-2 container">

            <div className="row g-5">

                {cardData.map((item, index) => (

                    <div className="col-md-3 " key={index}>
                        <div className="dashboard-card">
                            <div className="icon-box">
                                <i className={`${item.icon}`}></i>
                                <i></i>
                            </div>
                            <h2>{item.number}</h2>
                            <h5>{item.title}</h5>
                            <p>{item.subtitle}</p>
                        </div>

                    </div>

                ))}

            </div>

        </div>
    )
}