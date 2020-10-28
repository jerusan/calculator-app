import React, { Component } from "react";

class RecentOperationsComponent extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      List: []
    }
    this.grabRecentOperations()
    this.connect();
  }

  grabRecentOperations() {
    const requestOptions = {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'access-control-allow-origin': '*'
      }
    };
    fetch('http://localhost:3000/recentOperations', requestOptions)
    .then(response => response.json())
    .then(data => { 
      this.setState({
        List: data
      })});
  }

  // ToDo: Put ws initialization in a separate file
  connect() {
    var socket = new WebSocket("ws://localhost:8080/ws");
    console.log("Attempting Connection...");

    socket.onopen = () => {
      console.log("Successfully Connected");
    };

    socket.onmessage = msg => {
      this.update(msg);
    };

    socket.onclose = event => {
      console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
      console.log("Socket Error: ", error);
    };
  };

  update(data) {
    this.setState({
      List: JSON.parse(data.data)
    });
  }

  render() {
    return (
      <span>
        <h4>List of recently executed operations</h4>
        {this.state.List.map(item => (        
            <div> {item} </div>      
        ))}
      </span>)
  }
}

export default RecentOperationsComponent;