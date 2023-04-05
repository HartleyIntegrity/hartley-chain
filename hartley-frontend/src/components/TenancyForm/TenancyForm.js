import React, { useState } from "react";
import styles from "./TenancyForm.module.css";

const TenancyForm = ({ onSubmit }) => {
  const [formData, setFormData] = useState({
    landlordName: "",
    tenantName: "",
    startDate: "",
    endDate: "",
    rentAmount: "",
  });

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({
      ...prevFormData,
      [name]: value,
    }));
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    onSubmit(formData);
    setFormData({
      landlordName: "",
      tenantName: "",
      startDate: "",
      endDate: "",
      rentAmount: "",
    });
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <h2 className={styles.title}>Create New Tenancy Agreement</h2>
      <div className={styles.field}>
        <label htmlFor="landlordName">Landlord Name</label>
        <input
          type="text"
          id="landlordName"
          name="landlordName"
          value={formData.landlordName}
          onChange={handleInputChange}
          required
        />
      </div>
      <div className={styles.field}>
        <label htmlFor="tenantName">Tenant Name</label>
        <input
          type="text"
          id="tenantName"
          name="tenantName"
          value={formData.tenantName}
          onChange={handleInputChange}
          required
        />
      </div>
      <div className={styles.field}>
        <label htmlFor="startDate">Start Date</label>
        <input
          type="date"
          id="startDate"
          name="startDate"
          value={formData.startDate}
          onChange={handleInputChange}
          required
        />
      </div>
      <div className={styles.field}>
        <label htmlFor="endDate">End Date</label>
        <input
          type="date"
          id="endDate"
          name="endDate"
          value={formData.endDate}
          onChange={handleInputChange}
          required
        />
      </div>
      <div className={styles.field}>
        <label htmlFor="rentAmount">Rent Amount</label>
        <input
          type="number"
          id="rentAmount"
          name="rentAmount"
          value={formData.rentAmount}
          onChange={handleInputChange}
          required
        />
      </div>
      <button type="submit" className={styles.submitButton}>
        Create
      </button>
    </form>
  );
};

export default TenancyForm;
