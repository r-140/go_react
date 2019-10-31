import React from 'react';
import './App.css';
import { Route, BrowserRouter } from 'react-router-dom';
import Home from './components/layouts/Home';
import About from './components/layouts/About';
import Layout from './components/layouts/Layout';
import NewsArticle from './components/containers/NewsArticle';
import NewsSubmit from './components/containers/NewsSubmit';

import { Provider } from 'react-redux';
import store from './stores/store';

import { ApolloProvider } from 'react-apollo';
import client from './graphqlconfig/graphqlconfig' 

  const App = () => (
    <ApolloProvider client={client}>
        <Provider store={store}>
            <BrowserRouter>
                <Layout>
                    <Route exact path="/" component={Home } />
                    <Route path="/about" component={About} /> 
                    <Route path='/news/:id' component={NewsArticle}/>  
                    <Route path='/submit' component={NewsSubmit}/>  
                </Layout> 
            </BrowserRouter>
        </Provider> 
    </ApolloProvider>
  );
  
  export default App;