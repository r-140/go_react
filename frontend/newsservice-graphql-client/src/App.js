import React from 'react';
import logo from './logo.svg';
import ApolloClient from "apollo-boost";
import { ApolloProvider } from "react-apollo";

import News from './containers/components/News'
import './App.css';

const NEWSSERVICE_BASE_URL = `${process.env.REACT_APP_API_PROXY}/graphql`


const client = new ApolloClient({
  uri: NEWSSERVICE_BASE_URL
});

const App = () => (
  <ApolloProvider client={client}>
    <div>
      <News/>
    </div>
  </ApolloProvider>
);


// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }

export default App;
