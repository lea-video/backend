package model.core;

public class PlaylistItem {
    int index;
    Playable playable;
    int removedAt;
    int addedAt;

    public PlaylistItem(int idx, Playable m) {
        this.index = idx;
        this.playable = m;
    }
}
