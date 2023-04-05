import { useState, useEffect } from "react";
import TenancyForm from "./components/TenancyForm/TenancyForm";
import TenancyList from "./components/TenancyList/TenancyList";
import { getTenancies, createTenancy } from "./api";

function App() {
  const [tenancies, setTenancies] = useState([]);

  const fetchTenancies = async () => {
    const data = await getTenancies();
    setTenancies(data);
  };

  const handleCreateTenancy = async (tenancy) => {
    await createTenancy(tenancy);
    fetchTenancies();
  };

  useEffect(() => {
    fetchTenancies();
  }, []);

  return (
    <div className="container">
      <h1 className="title">Hartley Property Rentals</h1>
      <TenancyForm onCreateTenancy={handleCreateTenancy} />
      <TenancyList tenancies={tenancies} />
    </div>
  );
}

export default App;
