import React, { Component } from "react";
import { useParams } from "react-router-dom";

export default function Category(props) {
  let { category } = useParams();

  return <h2>{category} Movies</h2>;
}
