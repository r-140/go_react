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


// const CREATE_NEWS = gql`
//     mutation CreateNewsMutation {
//         CreateNewsMutation($title: title!, $teaser: teaser!, body: body!) {
//         id
//         title
//         teaser
//         body
//         }
//     }
// `;


// const ADD_COMMENT = gql`
//     mutation AddCommentMutation {
//         AddCommentMutation($newsID: ID!, $username: username!, $body: body!) {
//         username
//         body
//         }
//     }
// `;

export { GET_ALL_NEWS, GET_NEWS_BY_ID };

// const STAR_REPOSITORY = gql`
//   mutation($id: ID!) {
//     addStar(input: { starrableId: $id }) {
//       starrable {
//         id
//         viewerHasStarred
//       }
//     }
//   }
// `;