import { useRecoilState } from "recoil";
import { profileState } from "../recoil/profile.recoil";
import { useEffect, useState } from "react";
import { api } from "@/api/defaultApi";
import { v4 as uuidv4 } from "uuid";
import { Record } from "api/api";

const headers = {
    Authorization: "",
};
function useProfile() {
    const [profile, setProfile] = useRecoilState(profileState);
    const [records, setRecords] = useState<Record[]>();
    useEffect(() => {
        if (profile.nickname) {
            headers.Authorization = profile.nickname;
        } else {
            const uuid: string = uuidv4();
            // const name = btoa("JaeHwan"); // 대안.
            headers.Authorization = uuid;
        }
        api.getMe({ headers })
            .then((me) => {
                const getProfile = {
                    nickname: me.data.nickname,
                    elo_score: me.data.elo_score,
                };
                setProfile(getProfile);
            })
            .catch((error) => {
                console.log("getMe Error : ", error);
            });
        api.getRecords()
            .then((records) => {
                console.log("🚀 ~ file: Profile.hook.tsx:35 ~ .then ~ records:", records);
                setRecords(records.data.records);
            })
            .catch((error) => {
                console.log("getRecords Error : ", error);
            });
        return () => {};
    }, []);
    return { profile, records };
}

export default useProfile;
