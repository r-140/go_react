import React, { Component} from 'react';
import { GET_NEWS_BY_ID } from '../../queries/query'
import { Query } from "react-apollo";
import NewsItemDetail from '../presentation/NewsItemDetail';
import CommentsPanel from './CommentsPanel';

class NewsArticle extends Component {

    render(){
        const id = this.props.match.params.id;
        console.log("properties params ", id)
        return (
            <div>
                <Query
                    query={GET_NEWS_BY_ID}
                    variables = {{id}}
                    >
                    {({ loading, error, data }) => {
                        if (loading) return <p>Loading...</p>;
                        if (error) {
                            console.log("error ", error.Message)
                            return <p>error </p>;

                        } 

                        return <ul><div><NewsItemDetail data={data.News} />
                         <CommentsPanel comments={data.News.comments} id={data.News.id} />
                         </div> </ul>
                        
                    }}
                </Query>

                {/* <ul>
                    { !this.props.newsItemLoading ? <div><NewsItemDetail data={this.props.newsItem} />
                     <CommentsPanel comments={this.props.comments} id={this.props.newsItem._id} /></div> : <div>Loading</div>}
                </ul> */}
            </div>
        )
    }

    
    
}

 


export default NewsArticle