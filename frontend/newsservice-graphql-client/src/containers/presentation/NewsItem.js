import React from 'react';
const NewsItem = (props) => (
    <div key={props.item.id}>
        <p>{`${props.item.title} by ${props.item.teaser}`}</p>
    </div>
);
export default NewsItem;