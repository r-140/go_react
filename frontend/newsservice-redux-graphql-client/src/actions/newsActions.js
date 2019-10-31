import actionTypes from '../constants/actionTypes';
import {GET_ALL_NEWS, GET_NEWS_BY_ID, CREATE_NEWS, ADD_COMMENT} from './newsServiceQLQueries'
import client from '../graphqlconfig/graphqlconfig' 

function addComment(username, body){
    return {
        type: actionTypes.NEWS_ADDCOMMENT,
        username: username,
        body: body
    }
}

function newsItemReceived(newsItem){
    return {
        type: actionTypes.NEWSITEM_RECEIVED,
        newsItem: newsItem
    }
}

function newsReceived(news){
    console.log(" newsReceived ", news)
    return {
        type: actionTypes.NEWS_RECEIVED,
        news: news
    }
}

export function fetchNews(){
    return dispatch => {
        client.query({
            query: GET_ALL_NEWS
          }).then(data => {
            dispatch(newsReceived(data.data.AllNews))
          }).catch(e => { console.log(e)
          });
    }    
}

export function fetchNewsItem(id){
    return dispatch => {
        client.query({
            query: GET_NEWS_BY_ID,
            variables: {id},
          }).then(data => {
            dispatch(newsItemReceived(data.data.News))
          }).catch(e => { console.log(e)
          });
    }    
}

// function newsItemLoading(){
//     return {
//         type: actionTypes.NEWSITEM_LOADING
//     }
// }

export function submitNewsStory(data){

    const title = data.title
    const teaser = data.teaser;
    const body = data.body
    return dispatch => {
        client.mutate({
            mutation: CREATE_NEWS,
            variables: {title, teaser, body}
          }).then(data => {
            dispatch(newsItemReceived(data.data.News))
          }).catch(e => { console.log(e)
          });
    }    
}

export function submitComment(newsItemID, username, data){
    // console.log("submitComment newsItemID  ", newsItemID)
    // console.log("submitComment data  ", data.body)

    return dispatch => {
            const body = data.body
            client.mutate({
                mutation: ADD_COMMENT,
                variables: {newsID: newsItemID, username, body}
              }).then(data => {
                  console.log("submitComment data received ", data)
                dispatch(addComment(username, data.body))
              }).catch(e => { console.log(e)
              });
        }    
        // return fetch(`${process.env.REACT_APP_API_PROXY}/news/${newsItemID}/comment`, { 
        //     method: 'POST', 
        //      headers: {
        //         'Accept': 'application/json',
        //         'Content-Type': 'application/json',
        //       },
        //     body: JSON.stringify(data),
        //     mode: 'cors'
        // })
        //     .then( (response) => {
        //         if (!response.ok) {
        //             throw Error(response.statusText);
        //         }else{

        //             dispatch(addComment(username, data.body))
        //         }
        //     })
        //     .catch( (e) => console.log(e) );
    // }    
}