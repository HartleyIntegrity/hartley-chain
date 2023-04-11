import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';

const TenancyList = () => {
  const [tenancies, setTenancies] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8000/tenancies')
      .then(response => setTenancies(response.data))
      .catch(error => console.log(error));
  }, []);

  return (
    <div>
      <h1>Tenancy List</h1>
      <table className="table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Landlord Name</th>
            <th>Tenant Name</th>
            <th>Start Date</th>
            <th>End Date</th>
            <th>Rent Amount</th>
            <th>Deposit</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {tenancies && tenancies.map(tenancy => (
            <tr key={tenancy.id}>
              <td>{tenancy.id}</td>
              <td>{tenancy.landlord_name}</td>
              <td>{tenancy.tenant_name}</td>
              <td>{tenancy.start_date}</td>
              <td>{tenancy.end_date}</td>
              <td>{tenancy.rent_amount}</td>
              <td>{tenancy.deposit}</td>
              <td>
                <Link to={`/tenancies/${tenancy.id}`} className="btn btn-primary mr-2">View</Link>
                <Link to={`/tenancies/${tenancy.id}/edit`} className="btn btn-success mr-2">Edit</Link>
                <button className="btn btn-danger" onClick={() => {
                  axios.delete(`http://localhost:8000/tenancies/${tenancy.id}`)
                    .then(response => {
                      if (response.status === 204) {
                        setTenancies(tenancies.filter(t => t.id !== tenancy.id));
                      }
                    })
                    .catch(error => console.log(error));
                }}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <Link to="/create" className="btn btn-primary">Create New Tenancy</Link>
    </div>
  );
};

export default TenancyList;
