// src/components/TenancyList.js
import React, { Component } from 'react';
import axios from 'axios';

class TenancyList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      tenancies: [],
    };
  }

  componentDidMount() {
    // TODO: Replace "http://localhost:8080" with the actual API URL
    axios.get('http://localhost:8080/tenancies')
      .then(response => {
        this.setState({ tenancies: response.data });
      });
  }

  // src/components/TenancyList.js (update the render method)

    render() {
    const { tenancies } = this.state;
  
    return (
      <div>
        <h2>Tenancy List</h2>
        <ul>
          {tenancies.map(tenancy => (
            <li key={tenancy.ID}>
              {/* TODO: Replace with the link to the TenancyDetails component */}
              {tenancy.propertyID} - {tenancy.tenantID}
            </li>
          ))}
        </ul>
      </div>
    );
  }
  
}

export default TenancyList;
