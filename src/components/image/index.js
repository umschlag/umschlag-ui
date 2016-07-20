import React from 'react';
import './index.less';

export default class Image extends React.Component {
  render() {
    const {
      src,
      alt,
      ...otherProps
    } = this.props;

    return (
      <img src={src} alt={alt} {...otherProps}/>
    );
  }
}
