import React, { useEffect, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import axios from 'axios';

const TenancyDetail = () => {
  const { id } = useParams();
  const [tenancy, setTenancy] = useState(null);

  useEffect(() => {
    axios.get(`http://localhost:8000/tenancies/${id}`)
      .then(response => setTenancy(response.data))
      .catch(error => console.log(error));
  }, [id]);

  if (!tenancy) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>Tenancy Detail</h1>
      <table className="table">
        <tbody>
          <tr>
            <td>ID</td>
            <td>{tenancy.id}</td>
          </tr>
          <tr>
            <td>Landlord Name</td>
            <td>{tenancy.landlord_name}</td>
          </tr>
          <tr>
            <td>Tenant Name</td>
            <td>{tenancy.tenant_name}</td>
          </tr>
          <tr>
            <td>Property ID</td>
            <td>{tenancy.property_id}</td>
          </tr>
          <tr>
            <td>Start Date</td>
            <td>{tenancy.start_date}</td>
          </tr>
          <tr>
            <td>End Date</td>
            <td>{tenancy.end_date}</td>
          </tr>
          <tr>
            <td>Rent Amount</td>
            <td>{tenancy.rent_amount}</td>
          </tr>
          <tr>
            <td>Deposit</td>
            <td>{tenancy.deposit}</td>
          </tr>
        </tbody>
      </table>
      <Link to={`/tenancies/${id}/edit`} className="btn btn-success mr-2">Edit</Link>
      <Link to="/" className="btn btn-primary">Back to List</Link>
    </div>
  );
};

export default TenancyDetail;
