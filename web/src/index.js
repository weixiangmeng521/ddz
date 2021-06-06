import React from 'react';
import ReactDOM from 'react-dom';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import routes from "./router/router"
import { Socket } from "./socket/socket"
import 'bootstrap/dist/css/bootstrap.css';

Socket.connect();

function RouteWithSubRoutes(route) {
  return (
    <Route
      path={route.path}
      render={props => (
        // pass the sub-routes down to keep nesting
        <route.component {...props} routes={route.routes} />
      )}
    />
  );
}

ReactDOM.render(
  <React.StrictMode>
      <Router>
        <Switch>
          {routes.map((route, i) => (
            <RouteWithSubRoutes exact key={i} {...route} />
          ))}
        </Switch>
      </Router>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
