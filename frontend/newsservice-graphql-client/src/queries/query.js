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

// https://stackoverflow.com/questions/51522902/apollo-query-with-variable

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

// const query = gql`
//   query User($okta: string) {
//     User(okta: $okta){
//       id
//     }
//   }
// `;
export { GET_ALL_NEWS, GET_NEWS_BY_ID };

