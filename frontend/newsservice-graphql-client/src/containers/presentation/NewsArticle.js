import React, { Component} from 'react';
// import NewsItemDetail from '../presentation/NewsItemDetail';
// import CommentsPanel from './CommentsPanel';
// import { connect } from 'react-redux'
// import { fetchNewsItem } from '../../actions/newsActions'

class NewsArticle extends Component {


    // componentDidMount(){
    //     this.props.dispatch(fetchNewsItem(this.props.match.params.id));
    // }

    render(){
        console.log("properties id ", this.props.match.params.id)
        return (
            <div>
                <h2>News Story</h2>

                {/* <ul>
                    { !this.props.newsItemLoading ? <div><NewsItemDetail data={this.props.newsItem} /> <CommentsPanel comments={this.props.comments} id={this.props.newsItem._id} /></div> : <div>Loading</div>}
                </ul> */}
            </div>
        )
    }
}



export default NewsArticle