import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom';

const ShowMovieGenre = () => {
    let {id} = useParams();
    const [movies, setMovies] = useState([]);
    const [loaded, setLoaded] = useState(false);
    const [errorMessage, setErrorMessage] = useState(null);

    useEffect(() => {
        const fetchMovies = async () => {
          try {
              const result = await axios(`http://localhost:8000/api/genre/${id}/movies`);
              await setMovies(result.data.movies);
              setLoaded(true);
          } catch (error) {
              setErrorMessage(error.response.message);
          }
        };
        fetchMovies();
      }, [id]);

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
                    {Array.isArray(movies) ? (
                        <ul>
                            {movies.map((movie, index) => (
                                <li key={index}>
                                    <Link to={`/movies/${movie.id}`}>{movie.title}</Link>
                                </li>
                            ))}
                        </ul>
                    ) : (
                        <p>No movies found for this genre.</p>
                    )}
                </div>
            )}
        </>
    )
}

export default ShowMovieGenre