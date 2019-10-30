import React, { Component} from 'react';


class CommentElement extends Component {
    render(){
        return (
            <div>
                {/* <div><b>{this.props.data.username}</b></div> */}
                <div>{this.props.data.body}</div>
            </div>
        )
    }
}




export default CommentElement