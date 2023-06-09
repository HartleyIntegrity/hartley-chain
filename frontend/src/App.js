import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import TenancyList from './TenancyList';
import TenancyCreate from './TenancyCreate';
import TenancyEdit from './TenancyEdit';
import TenancyDetail from './TenancyDetail';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<TenancyList />} />
        <Route exact path="/create" element={<TenancyCreate />} />
        <Route exact path="/edit/:id" element={<TenancyEdit />} />
        <Route exact path="/tenancies/:id/edit" element={<TenancyEdit />} />
        <Route exact path="/tenancies/:id" element={<TenancyDetail />} />
      </Routes>
    </Router>

  );
};

export default App;
