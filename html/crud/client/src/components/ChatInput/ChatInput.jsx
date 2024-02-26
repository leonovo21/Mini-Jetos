import React, { Component } from 'react';
import './ChatInput.scss';

class ChatInput extends Component {
  
  render() {
    return (
      <div className='ChatInput'>
        <textarea name='postcontent' onKeyDown={this.props.send} placeholder='Write here...'></textarea>
      </div>
    );
  };

}

export default ChatInput;
