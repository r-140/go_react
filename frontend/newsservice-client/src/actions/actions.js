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

export function fetchNews(){
    return dispatch => {
        return fetch(`${process.env.REACT_APP_API_PROXY}/news`)
        .then( (response) => response.json() )
        .then( (data) => {
            dispatch(newsReceived(data))
        })
        .catch( (e) => console.log(e) );
    }    
}

export function fetchNewsItem(id){
    return dispatch => {
        return fetch(`${process.env.REACT_APP_API_PROXY}/news/${id}`)
        .then( (response) => response.json() )
        .then( (data) => dispatch(newsItemReceived(data)))
        .catch( (e) => console.log(e) );
    }    
}

function newsItemLoading(){
    return {
        type: actionTypes.NEWSITEM_LOADING
    }
}

export function submitNewsStory(data){
    return dispatch => {
        return fetch(`${process.env.REACT_APP_API_PROXY}/news`, { 
            method: 'POST', 
             headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
              },
            body: JSON.stringify(data), 
            mode: 'cors'})
            .catch( (e) => console.log(e) );
    }    
}