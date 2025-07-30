package worker;

import redis.clients.jedis.Jedis;
import redis.clients.jedis.exceptions.JedisConnectionException;
import java.sql.*;
import org.json.JSONObject;

public class Worker {
    public static void main(String[] args) {
        try (Jedis redis = connectToRedis("redis");
             Connection dbConn = connectToDB("db")) {
            System.err.println("Watching vote queue");

            while (true) {
                String voteJSON = redis.blpop(0, "votes").get(1);
                JSONObject voteData = new JSONObject(voteJSON);
                String voterID = voteData.getString("voter_id");
                String vote = voteData.getString("vote");

                System.err.printf("Processing vote for '%s' by '%s'%n", vote, voterID);
                updateVote(dbConn, voterID, vote);
            }
        } catch (Exception e) {
            e.printStackTrace();
            System.exit(1);
        }
    }

    static void updateVote(Connection dbConn, String voterID, String vote) throws SQLException {
        try (PreparedStatement insert = dbConn.prepareStatement(
                "INSERT INTO votes (id, vote) VALUES (?, ?)")) {
            insert.setString(1, voterID);
            insert.setString(2, vote);
            insert.executeUpdate();
        } catch (SQLException e) {
            try (PreparedStatement update = dbConn.prepareStatement(
                    "UPDATE votes SET vote = ? WHERE id = ?")) {
                update.setString(1, vote);
                update.setString(2, voterID);
                update.executeUpdate();
            }
        }
    }

    static Jedis connectToRedis(String host) {
        String redisHost = System.getenv("REDIS_HOST") != null ? System.getenv("REDIS_HOST") : host;
        int redisPort = System.getenv("REDIS_PORT") != null ? Integer.parseInt(System.getenv("REDIS_PORT")) : 6379;
        Jedis conn = new Jedis(redisHost, redisPort);

        while (true) {
            try {
                conn.ping();
                System.err.println("Connected to Redis at " + redisHost + ":" + redisPort);
                break;
            } catch (JedisConnectionException e) {
                System.err.println("Waiting for Redis at " + redisHost + ":" + redisPort + ": " + e.getMessage());
                sleep(1000);
            }
        }

        return conn;
    }

    static Connection connectToDB(String host) throws SQLException {
        Connection conn = null;
        try {
            Class.forName("org.postgresql.Driver");
            String dbHost = System.getenv("DB_HOST") != null ? System.getenv("DB_HOST") : host;
            String dbUser = System.getenv("DB_USER") != null ? System.getenv("DB_USER") : "postgres1";
            String dbPassword = System.getenv("DB_PASSWORD") != null ? System.getenv("DB_PASSWORD") : "postgres1";
            String dbName = System.getenv("DB_NAME") != null ? System.getenv("DB_NAME") : "postgres1";
            String url = "jdbc:postgresql://" + dbHost + "/" + dbName;

            while (conn == null) {
                try {
                    conn = DriverManager.getConnection(url, dbUser, dbPassword);
                    System.err.println("Connected to DB at " + url + " with user " + dbUser);
                } catch (SQLException e) {
                    System.err.println("Waiting for DB at " + url + " with user " + dbUser + ": " + e.getMessage());
                    sleep(1000);
                }
            }

            try (PreparedStatement st = conn.prepareStatement(
                    "CREATE TABLE IF NOT EXISTS votes (id VARCHAR(255) NOT NULL UNIQUE, vote VARCHAR(255) NOT NULL)")) {
                st.executeUpdate();
            }
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
            System.exit(1);
        }
        return conn;
    }

    static void sleep(long duration) {
        try {
            Thread.sleep(duration);
        } catch (InterruptedException e) {
            System.exit(1);
        }
    }
}