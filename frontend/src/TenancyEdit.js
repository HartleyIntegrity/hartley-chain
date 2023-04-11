import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import axios from 'axios';

const TenancyEdit = () => {
  const [landlordName, setLandlordName] = useState('');
  const [tenantName, setTenantName] = useState('');
  const [propertyId, setPropertyId] = useState('');
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const [rentAmount, setRentAmount] = useState('');
  const [deposit, setDeposit] = useState('');
  const navigate = useNavigate();
  const { id } = useParams();

  useEffect(() => {
    axios.get(`http://localhost:8000/tenancies/${id}`)
      .then(response => {
        const tenancy = response.data;
        setLandlordName(tenancy.landlord_name);
        setTenantName(tenancy.tenant_name);
        setPropertyId(tenancy.property_id);
        setStartDate(tenancy.start_date);
        setEndDate(tenancy.end_date);
        setRentAmount(tenancy.rent_amount);
        setDeposit(tenancy.deposit);
      })
      .catch(error => console.log(error));
  }, [id]);

  const handleSubmit = e => {
    e.preventDefault();
    axios.put(`http://localhost:8000/tenancies/${id}`, {
      landlord_name: landlordName,
      tenant_name: tenantName,
      property_id: propertyId,
      start_date: startDate,
      end_date: endDate,
      rent_amount: rentAmount,
      deposit: deposit
    })
      .then(() => {
        navigate.push('/');
      })
      .catch(error => console.log(error));
  };

  return (
    <div>
      <h1>Edit Tenancy</h1>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="landlordName">Landlord Name</label>
          <input type="text" className="form-control" id="landlordName" value={landlordName} onChange={e => setLandlordName(e.target.value)} />
        </div>
        <div className="form-group">
          <label htmlFor="tenantName">Tenant Name</label>
          <input type="text" className="form-control" id="tenantName" value={tenantName} onChange={e => setTenantName(e.target.value)} />
        </div>
        <div className="form-group">
          <label htmlFor="propertyId">Property ID</label>
          <input type="text" className="form-control" id="propertyId" value={propertyId} onChange={e => setPropertyId(e.target.value)} />
        </div>
        <div className="form-group">
          <label htmlFor="startDate">Start Date</label>
          <input type="text" className="form-control" id="startDate" value={startDate} onChange={e => setStartDate(e.target.value)} />
        </div>
        <div className="form-group">
          <label htmlFor="endDate">End Date</label>
          <input type="text" className="form-control" id="endDate" value={endDate} onChange={e => setEndDate(e.target.value)} />
        </div>
        <div className="form-group">
          <label htmlFor="rentAmount">Rent Amount</label>
          <input type="text" className="form-control" id="rentAmount" value={rentAmount} onChange={e => setRentAmount(e.target.value)} />
        </div>
        <div className="form-group">
          <label htmlFor="deposit">Deposit</label>
<input type="text" className="form-control" id="deposit" value={deposit} onChange={e => setDeposit(e.target.value)} />
</div>
<button type="submit" className="btn btn-primary mr-2">Update</button>
<button type="button" className="btn btn-secondary" onClick={() => navigate.push('/')}>Cancel</button>
</form>
</div>
);
};

export default TenancyEdit;