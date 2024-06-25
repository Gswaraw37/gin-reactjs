import React, { useEffect, useState } from 'react'
import MovieDetail from '../../components/movies/MovieDetail'
import { useParams } from 'react-router-dom';
import axios from 'axios';

const ShowMovie = () => {
    let {id} = useParams();
    const [movie, setMovie] = useState([]);
    const [loaded, setLoaded] = useState(false);
    const [errorMessage, setErrorMessage] = useState(null);

    useEffect(() => {
        const fetchMovies = async () => {
          try {
              const result = await axios(`http://localhost:8000/api/movies/${id}`);
              await setMovie(result.data.data);
              setLoaded(true);
          } catch (err) {
              setErrorMessage(err.response.message);
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
                <MovieDetail movie={movie} />
            )}
        </>
    )
}

export default ShowMovie