import actionTypes from '../constants/actionTypes';

function newsReceived(news){
    return {
        type: actionTypes.NEWS_RECEIVED,
        news: news
    }
}