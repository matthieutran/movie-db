import React, { Component, Fragment } from "react";

export default class Movie extends Component {
  state = { movie: {} };

  componentDidMount() {
    const id = this.props.match.params.id;
    this.setState({
      movie: {
        id: id,
        title: "Some Movie",
        runtime: 150,
      },
    });
  }

  render() {
    return (
      <Fragment>
        <h2>
          {this.state.movie.title} ({this.state.movie.id})
        </h2>
        <table className="table table-compact table-striped">
          <thead></thead>
          <tbody>
            <tr>
              <td>
                <strong>Title:</strong>
              </td>
              <td>{this.state.movie.title}</td>
            </tr>
            <tr>
              <td>
                <strong>Runtime:</strong>
              </td>
              <td>{this.state.movie.runtime} minutes</td>
            </tr>
          </tbody>
        </table>
      </Fragment>
    );
  }
}
