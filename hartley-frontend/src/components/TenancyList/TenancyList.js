import React from "react";
import styles from "./TenancyList.module.css";

const TenancyList = ({ tenancies }) => {
  return (
    <div className={styles.container}>
      <h2 className={styles.title}>Tenancy Agreements</h2>
      <table className={styles.table}>
        <thead>
          <tr>
            <th>Landlord Name</th>
            <th>Tenant Name</th>
            <th>Start Date</th>
            <th>End Date</th>
            <th>Rent Amount</th>
          </tr>
        </thead>
        <tbody>
          {tenancies.map((tenancy) => (
            <tr key={tenancy.id}>
              <td>{tenancy.landlordName}</td>
              <td>{tenancy.tenantName}</td>
              <td>{tenancy.startDate}</td>
              <td>{tenancy.endDate}</td>
              <td>£{tenancy.rentAmount.toFixed(2)}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TenancyList;
