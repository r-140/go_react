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


export { GET_ALL_NEWS };