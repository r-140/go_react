
import React from 'react';
import NewsItemListing from '../presentation/NewsItemListing';
const NewsList = ({
    news,
    
  }) => (
      <ul>
          {news.map( (item, i) => {
              return <li key={i}><NewsItemListing data = {item} /></li> 
        })}
      </ul>
    
    );
 
  export default NewsList


  