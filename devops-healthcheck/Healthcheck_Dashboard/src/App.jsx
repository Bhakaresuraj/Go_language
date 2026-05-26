
import './App.css'
import Sidebar from './components/sidebar/sidebar'
import Card from './components/Dashboard/cards'
import Navbar from './components/navbar/navbar'
function App() {

  return (
    <div className="d-flex">
      <Sidebar />
      <div className=" w-100">  
        <Navbar></Navbar>
        <Card></Card>
      </div>

    </div>


  )
}

export default App
