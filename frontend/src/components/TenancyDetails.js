// src/components/TenancyDetails.js
import React, { Component } from 'react';
import axios from 'axios';

class TenancyDetails extends Component {
  constructor(props) {
    super(props);
    this.state = {
      tenancy: {},
    };
  }

  componentDidMount() {
    const { match } = this.props;
    // TODO: Replace "http://localhost:8080" with the actual API URL
    axios.get(`http://localhost:8080/tenancies/${match.params.id}`)
      .then(response => {
        this.setState({ tenancy: response.data });
      });
  }

  // src/components/TenancyDetails.js (update the render method)

render() {
    const { tenancy } = this.state;
  
    return (
      <div>
        <h2>Tenancy Details</h2>
        <dl>
          <dt>Property ID</dt>
          <dd>{tenancy.propertyID}</dd>
          <dt>Tenant ID</dt>
          <dd>{tenancy.tenantID}</dd>
          <dt>Start Date</dt>
          <dd>{tenancy.startDate}</dd>
          <dt>End Date</dt>
          <dd>{tenancy.endDate}</dd>
        </dl>
      </div>
    );
  }
  
}

export default TenancyDetails;
