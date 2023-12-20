import React, { Component, useState } from 'react';
import ReactDOM from 'react-dom';
import AppFooter from './AppFooter';
import './index.css'



import AppContent from './AppContent';
import AppHeader from './AppHeader';
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.bundle.min.js'
console.log('Im running')




class App extends Component{
  render(){
    const  myProps = {
      title: "New app",
      subject: "New subject",
      favorit_lang: "Go",
    }
    return ( 
      <body>
        <style>{'body {background-color: wheat; }'}</style>
          <AppHeader {...myProps}/>
          <AppContent />
        <AppFooter>
        </AppFooter>
      </body>
    );
  }
}
ReactDOM.render(<App />, document.getElementById("root"));