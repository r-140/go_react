import React from 'react';
import { Link } from 'react-router-dom';

const NewsItem = (props) => (
    <div key={props.item.id}>
        <div><Link to={`/news/${props.item.id}`}><b>{props.item.title}</b></Link></div> 
    </div>
);
export default NewsItem;