import actionTypes from '../constants/actionTypes';

function newsItemReceived(newsItem){
    return {
        type: actionTypes.NEWSITEM_RECEIVED,
        newsItem: newsItem
    }
}

function newsReceived(news){
    return {
        type: actionTypes.NEWS_RECEIVED,
        news: news
    }
}

export function fetchNews(fakeNews){
    console.log('presend')
    return dispatch => {
        return fetch(`http://localhost:6768/news`)
        .then( (response) =>{
            console.log(response);
        });
    }    
}

export function fetchNewsItem(fakeNewsItem){
    return dispatch => {
        dispatch(newsItemReceived(fakeNewsItem));
    }
}