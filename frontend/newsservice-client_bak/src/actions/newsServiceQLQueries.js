import gql from 'graphql-tag';


// const GET_NEWS_BY_ID = gql`
//     {
//         News($id: ID!)) {
//           title
//           teaser
//           body
//           comments {
//             username
//             body
//           }
//         }
//       }
// `;

const  GET_ALL_NEWS = gql`
{
    AllNews {
        id
        title
        teaser
        
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

export { GET_ALL_NEWS };

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