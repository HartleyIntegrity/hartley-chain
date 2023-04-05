import React, { useState, useEffect } from 'react';
import PropertyList from './components/PropertyList';
import PropertyForm from './components/PropertyForm';
import Api from './api';

function App() {
  const [properties, setProperties] = useState([]);

  useEffect(() => {
    Api.getProperties().then(setProperties);
  }, []);

  const handleAddProperty = (property) => {
    Api.addProperty(property).then(() => {
      setProperties([...properties, property]);
    });
  };

  return (
    <div className="App">
      <h1>Hartley Property Rental</h1>
      <PropertyForm onAddProperty={handleAddProperty} />
      <PropertyList properties={properties} />
    </div>
  );
}

export default App;
