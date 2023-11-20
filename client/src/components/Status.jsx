import React from "react";
import clsx from 'clsx'
import './Status.css'

const Status = ({ url, statusCode }) => {
  return <li className="status">
    <span className="url">{url}</span>
    <i className={clsx('checker', {
      ok: statusCode === 200
    })} />
  </li>
}

export default Status;