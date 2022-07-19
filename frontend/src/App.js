import React, { Fragment } from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import Admin from "./pages/Admin";
import Category from "./pages/Category";
import Categories from "./pages/Categories";
import Home from "./pages/Home";
import Movie from "./pages/Movie";
import Movies from "./pages/Movies";

export default function App() {
  return (
    <Router>
      <div className="container">
        <div className="row">
          <h1 className="mt-3">Movie Database</h1>
        </div>

        <hr className="mb-3" />

        <div className="row">
          <div className="col-md-2">
            <nav>
              <ul className="list-group">
                <li className="list-group-item">
                  <Link exact to="/">
                    Home
                  </Link>
                </li>
                <li className="list-group-item">
                  <Link to="/movies">Movies</Link>
                </li>
                <li className="list-group-item">
                  <Link to="/categories">Categories</Link>
                </li>
                <li className="list-group-item">
                  <Link to="/admin">Manage Catalogue</Link>
                </li>
              </ul>
            </nav>
          </div>

          <div className="col-md-10">
            <Switch>
              <Route path="/movies/:id" component={Movie} />
              <Route path="/movies">
                <Movies />
              </Route>
              <Route exact path="/categories">
                <Categories />
              </Route>
              <Route
                path="/categories/:category"
                render={(props) => <Category {...props} />}
              ></Route>
              <Route path="/admin">
                <Admin />
              </Route>
              <Route path="/">
                <Home />
              </Route>
            </Switch>
          </div>
        </div>
      </div>
    </Router>
  );
}
