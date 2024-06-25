import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom';
import axios from 'axios';

const MovieList = () => {
    const [movies, setMovies] = useState([]);
    const [loaded, setLoaded] = useState(false);
    const [errorMessage, setErrorMessage] = useState(null);

    useEffect(() => {
      const fetchMovies = async () => {
        try {
            const result = await axios(`http://localhost:8000/api/movies/`);
            await setMovies(result.data.data);
            setLoaded(true);
        } catch (err) {
            setErrorMessage(err.response.message);
        }
      };
      fetchMovies();
    }, []);
    

    return (
        <>
            { !loaded ? (
                (() => {
                    if (errorMessage) {
                        return(
                            <div className="row">
                                <p>Opss... {errorMessage}</p>
                            </div>
                        );
                    } else {
                        return(
                            <div className="row">
                                <p>Loading...</p>
                            </div>
                        );
                    }
                })()
            ) : (
                <div className="row">
                    {movies.map((data, index) => (
                        <div className="col-sm-4 mb-2" key={index}>
                            <div className="card">
                            <div className="card-body">
                                <h5 className="card-title">{data.title}</h5>
                                <p className="card-text">{data.description}</p>
                                <Link to={`/movies/${data.id}`} className="btn btn-primary">
                                    Detail
                                </Link>
                            </div>
                            </div>
                        </div>
                    ))}
                </div>
            )}
        </>
    )
}

export default MovieList