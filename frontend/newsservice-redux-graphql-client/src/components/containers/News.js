import React from 'react';

import { Query } from 'react-apollo';
import NewsList from './NewsList'

import ErrorMessage from '../../Error';

import { connect } from 'react-redux'

import { GET_ALL_NEWS } from '../../actions/newsServiceQLQueries'

const News = () => (
    <Query
    query={GET_ALL_NEWS}
      notifyOnNetworkStatusChange={true}
    //   pollInterval={500}
    >
      {
          ({ data: { AllNews }, error }) => {
            console.log("news list ", AllNews)

            const { news } = AllNews;
    
            if (error) {
                return <ErrorMessage error={error} />;
            }
    
            return (
                <NewsList news={news}/>
            );
        }
      }
    </Query>
  );

  const mapStateToProps = state => {
    return {
        news: state.news.news
    }
}

export default connect(mapStateToProps)(News)

