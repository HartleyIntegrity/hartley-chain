import React, { useState, useEffect } from "react";
import TenancyForm from "../TenancyForm/TenancyForm";
import TenancyList from "../TenancyList/TenancyList";
import styles from "./App.module.css";

const App = () => {
  const [tenancies, setTenancies] = useState([]);

  useEffect(() => {
    // TODO: Fetch tenancies from the API and set the state
  }, []);

  const handleTenancySubmit = (formData) => {
    // TODO: Submit the formData to the API and update the tenancies state
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Hartley Property Rental</h1>
      <TenancyForm onSubmit={handleTenancySubmit} />
      <TenancyList tenancies={tenancies} />
    </div>
  );
};

export default App;
