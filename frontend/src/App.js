// src/App.js
import React, { Component } from 'react';
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom";
import './App.css';

// TODO: Add components for TenancyList, TenancyForm, and TenancyDetails
import TenancyList from './components/TenancyList';
import TenancyForm from './components/TenancyForm';
import TenancyDetails from './components/TenancyDetails';


class App extends Component {
  render() {
    return (
      <BrowserRouter>
      <Routes>
          <Route exact path="/" component={TenancyList} />
          <Route path="/create" component={TenancyForm} />
          <Route path="/tenancies/:id" component={TenancyDetails} />
      </Routes>
      </BrowserRouter>
    );
  }
}

export default App;
