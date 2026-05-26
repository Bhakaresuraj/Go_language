import './App.css'
import { GetAllServices } from './utils/api';
import { useRef, useEffect, useState } from 'react';
// import { } from "react";
import Sidebar from './components/sidebar/sidebar'
import Card from './components/Dashboard/cards'
import Navbar from './components/navbar/navbar'
import ServiceTable from './components/Service-Table/servicetable'
import AddServiceModal from './components/model/AddserviceModal'
function App() {

  const [showModal, setShowModal] = useState(false);
  const goToTop = () => {
    topRef.current.scrollIntoView({
      behavior: "smooth"
    });
  };
  const goToServices = () => {
    tableRef.current.scrollIntoView({
      behavior: "smooth"
    });
  };
  const topRef = useRef(null);
  const tableRef = useRef(null);
  const [services, setServices] = useState([]);
  async function fetchServices() {
    try {
      let data = await GetAllServices();
      setServices(data);
    } catch (err) {
      console.log(err);
    }
  }
  useEffect(() => {

    fetchServices();
    const interval = setInterval(() => {

      fetchServices();

    }, 10000);
    return () => clearInterval(interval);
  }, []);
  return (
    <div className="d-flex Container ">
      <Sidebar openModal={() => setShowModal(true)} goToTop={goToTop} goToServices={goToServices} />
      <div className=" content ">

        <div ref={topRef}>
          <Navbar></Navbar>
        </div>
        <Card data={services}></Card>
        <div ref={tableRef}>
          <ServiceTable data={services}></ServiceTable>
        </div>
      </div>
      <AddServiceModal
        show={showModal}
        onClose={() =>
          setShowModal(false)
        }
        refreshServices={fetchServices}
      />

    </div>
  )
}
export default App
