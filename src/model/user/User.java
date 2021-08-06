package model.user;

import model.utility.Celebrity;
import model.core.Playlist;
import model.core.Snippet;

public class User {
    String username;
    String displayName;
    Credentials[] credentials;
    Playlist[] playlists;
    Snippet[] snippets;
    Celebrity[] follows;
}
