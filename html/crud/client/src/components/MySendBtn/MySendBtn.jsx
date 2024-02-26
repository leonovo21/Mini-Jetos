import React,{Component} from "react";
import './MySendBtn.scss'

class MySendBtn extends Component{
    render(){
        return(
            <div className="MSBT">
                <textarea name="mscontent" onKeyDown={this.props.send} placeholder="I did this"></textarea>
            </div>
        )
    }
}
export default MySendBtn