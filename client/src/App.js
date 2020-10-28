import React, { Component } from "react";
import CalculatorComponent from './Components/Calculator/Calculator'
import RecentOperationsComponent from './Components/RecentOperations/RecentOperations'

class App extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="container">
        <h1>Simple Web Calculator</h1>
        <h4>only characters('+', '-', '/', '*') are allowed in operator field </h4>
        <h4>only numbers are allowed in operand fields</h4>
        <div className="children">
          <span className="calc"> <CalculatorComponent />  </span>
          <span className="list"> <RecentOperationsComponent /> </span>
        </div>
      </div>
    )
  }
}

export default App
