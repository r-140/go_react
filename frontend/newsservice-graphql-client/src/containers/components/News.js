import React from 'react';
import { Query } from "react-apollo";

import { GET_ALL_NEWS } from '../../queries/query'
import NewsItem from '../presentation/NewsItem';

const News = () => (
  <Query
    query={GET_ALL_NEWS}
  >
    {({ loading, error, data }) => {
      if (loading) return <p>Loading...</p>;
      if (error) return <p>Error :(</p>;
      return data.AllNews.map((newsItem) => (
        <NewsItem item={newsItem} />
      ));
    }}
  </Query>
);
export default News;