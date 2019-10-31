import { ApolloClient } from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';

 
const NEWSSERVICE_BASE_URL = `${process.env.REACT_APP_API_PROXY}/graphql`
const httpLink = new HttpLink({
  uri: NEWSSERVICE_BASE_URL,

});

const cache = new InMemoryCache();

const client = new ApolloClient({
    link: httpLink,
    cache,
  });

export default client