import gql from 'graphql-tag';


const  GET_ALL_NEWS = gql` 
{
    AllNews {
        id
        title
        teaser
        
      }
  }
`;

const GET_NEWS_BY_ID = gql` query GET_NEWS_BY_ID($id: String)
{
    News(id: $id) {
      id
      title
      teaser
      body
      comments {
        username
        body
      }
    }
  }
`;


const CREATE_NEWS = gql` mutation CREATE_NEWS ($title: String, $teaser: String, $body: String)
    {
      CreateNewsMutation(title: $title, teaser: $teaser, body: $body)
      {
        id
        title
        teaser
        body
        comments {
          body
        }
      }
    }
`;


const ADD_COMMENT = gql` mutation ADD_COMMENT ($newsID: String, $username: String, $body: String)
  {
      AddCommentMutation(newsID: $newsID, username: $username, body: $body) {
      username
      body
    }
  }
`;

export { GET_ALL_NEWS, GET_NEWS_BY_ID, CREATE_NEWS, ADD_COMMENT };

