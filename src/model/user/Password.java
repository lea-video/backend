package model.user;

public class Password implements Credentials {
    String hash;
    String salt;

    @Override
    public Session getSession() {
        return null;
    }

    @Override
    public boolean isValid() {
        return false;
    }
}
