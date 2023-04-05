// src/components/TenancyForm.js
import React, { Component } from 'react';
import axios from 'axios';

class TenancyForm extends Component {
  // src/components/TenancyForm.js (update the constructor, handleChange, and handleSubmit methods)

constructor(props) {
    super(props);
    this.state = {
      propertyID: '',
      tenantID: '',
      startDate: '',
      endDate: '',
    };
  
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  
  handleChange(event) {
    const { name, value } = event.target;
    this.setState({ [name]: value });
  }
  
  handleSubmit(event) {
    event.preventDefault();
    const { propertyID, tenantID, startDate, endDate } = this.state;
  
    // TODO: Replace "http://localhost:8080" with the actual API URL
    axios.post('http://localhost:8080/tenancies', { propertyID, tenantID, startDate, endDate })
      .then(response => {
        // TODO: Redirect to the TenancyDetails component
      });
  }
  

  render() {
    return (
      <div>
        <h2>Tenancy Form</h2>
        {/* TODO: Render the form */}
      </div>
    );
  }
}

export default TenancyForm;
