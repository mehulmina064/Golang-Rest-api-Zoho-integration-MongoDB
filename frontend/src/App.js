// import logo from './logo.svg';
// import './App.css';
import React from 'react';
import 'bootstrap/dist/css/bootstrap.css';

// import Orders from './components/orders/orders.components';
import Users from './components/users/users.components';
require('dotenv').config() 


function App() {
  return (

   <div> 
   <Users />
  {/* <Orders /> */}

 </div>
  );
}
  
export default App;
