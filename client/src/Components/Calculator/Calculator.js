import React, { Component } from "react";


class CalculatorComponent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      Result: '',
      operator: '',
      operand1: '',
      operand2: ''
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({ value: event.target.value });
  }

  handleSubmit(event) {
    const requestOptions = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'access-control-allow-origin': '*'
      },
      body: JSON.stringify({
        operand1: parseInt(this.state.operand1),
        operator: this.state.operator,
        operand2: parseInt(this.state.operand2)
      })
    };
    fetch('http://localhost:3000/', requestOptions)
      .then(response => response.json())
      .then(data => {
        this.setState({
          Result: data
        });

      });
    event.preventDefault();
  }

  render() {
    return (
      <span>
        <form onSubmit={this.handleSubmit}>
          <label>
            Operand1:
            <input type="number" default={null} pattern="[0-9]*" value={this.state.operand1} onChange={e => this.setState({ operand1: e.target.value })} />
          </label>
          <label>
            Operator:
            <input type="text" pattern="[+, *, /, -]" value={this.state.operator} onChange={e => this.setState({ operator: e.target.value })} />
          </label>
          <label>
            Operand2:
            <input type="number" default={null} pattern="[0-9]*" value={this.state.operand2} onChange={e => this.setState({ operand2: e.target.value })} />
          </label>
          <input type="submit" value="Submit" />
        </form>
        <div id="result">Result: {this.state.Result}</div>
      </span>
    );
  }
}

export default CalculatorComponent;