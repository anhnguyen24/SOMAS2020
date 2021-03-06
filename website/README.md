# SOMAS2020 Visualisation Website

## Getting started 

### Requirements
- [Node.js 14.x.x](https://nodejs.org/en/)
    - `node -v` should produce `v14.x.x`
- [Yarn 1.22.x](https://yarnpkg.com/getting-started/install)
    - `yarn --version` should produce `1.22.x`

### Install dependencies
`yarn install`

### Make sure you have the output from the Go program
The `output` folder in the root of the repository must be present. 

## Scripts

### `yarn start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

<!-- ### `yarn test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information. -->

### `yarn build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.

<!-- ### `yarn deploy`

Deploy the app into [GitHub Pages](https://somas2020.github.com/SOMAS2020).\
This should be run automatically by CI. -->

## Information
- This website uses [React](https://reactjs.org/) and [TypeScript](https://www.typescriptlang.org/).
- The library used for UI/UX is [React Bootstrap](https://react-bootstrap.github.io/).
- The graphing library used is [d3.js](https://d3js.org/). 
    - You do not strictly need to use d3: You can use any language to produce visualisations, then include them
    in the website.