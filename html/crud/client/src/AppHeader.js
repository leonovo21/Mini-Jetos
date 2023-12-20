import React, {Component, Fragment} from "react";
import './AppHeader.css'

export default class AppHeader extends Component{

    render(){
        return(
                <h1 className="title">{this.props.title}</h1>
        );
    }
}