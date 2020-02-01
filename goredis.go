package redis_namespace

import (
	"github.com/go-redis/redis/v7"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

type Client struct {
	GoRedis
	client    GoRedis
	namespace string
}

type GoRedis interface {
	redis.Cmdable
}

func NewGoRedisWithNamespace(namespace string, client GoRedis) *Client {
	return &Client{client, client, namespace + ":"}
}

func (c *Client) Del(keys ...string) *redis.IntCmd {
	var wrappedKeys []string

	for _, key := range keys {
		wrappedKeys = append(wrappedKeys, c.namespace+key)
	}
	return c.client.Del(wrappedKeys...)
}

func (c *Client) Unlink(keys ...string) *redis.IntCmd {
	var wrappedKeys []string

	for _, key := range keys {
		wrappedKeys = append(wrappedKeys, c.namespace+key)
	}
	return c.client.Unlink(wrappedKeys...)
}

func (c *Client) Dump(key string) *redis.StringCmd {
	return c.client.Dump(c.namespace + key)
}

func (c *Client) Exists(keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) Keys(pattern string) *redis.StringSliceCmd {
	var unwrappedKeys []string
	keys := c.client.Keys(c.namespace + pattern)

	for _, key := range keys.Val() {
		unwrappedKeys = append(unwrappedKeys, strings.TrimPrefix(key, c.namespace))
	}

	pointerVal := reflect.ValueOf(keys)
	val := reflect.Indirect(pointerVal)
	member := val.FieldByName("val")
	ptrToY := unsafe.Pointer(member.UnsafeAddr())
	realPtrToY := (*[]string)(ptrToY)
	*realPtrToY = unwrappedKeys

	return keys
}

func (c *Client) Migrate(host, port, key string, db int64, timeout time.Duration) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) Move(key string, db int64) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) ObjectRefCount(key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ObjectEncoding(key string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) ObjectIdleTime(key string) *redis.DurationCmd {
	panic("implement me")
}

func (c *Client) Persist(key string) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) PTTL(key string) *redis.DurationCmd {
	panic("implement me")
}

func (c *Client) RandomKey() *redis.StringCmd {
	panic("implement me")
}

func (c *Client) Rename(key, newkey string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) RenameNX(key, newkey string) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) Restore(key string, ttl time.Duration, value string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) RestoreReplace(key string, ttl time.Duration, value string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) Sort(key string, sort *redis.Sort) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) SortStore(key, store string, sort *redis.Sort) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) SortInterfaces(key string, sort *redis.Sort) *redis.SliceCmd {
	panic("implement me")
}

func (c *Client) Touch(keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) TTL(key string) *redis.DurationCmd {
	panic("implement me")
}

func (c *Client) Type(key string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *Client) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *Client) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *Client) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	panic("implement me")
}

func (c *Client) Append(key, value string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) BitCount(key string, bitCount *redis.BitCount) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) BitOpAnd(destKey string, keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) BitOpOr(destKey string, keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) BitOpXor(destKey string, keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) BitOpNot(destKey string, key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) BitPos(key string, bit int64, pos ...int64) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) Decr(key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) DecrBy(key string, decrement int64) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) Get(key string) *redis.StringCmd {
	return c.client.Get(c.namespace + key)
}

func (c *Client) GetBit(key string, offset int64) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) GetRange(key string, start, end int64) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) GetSet(key string, value interface{}) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) Incr(key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) IncrBy(key string, value int64) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) IncrByFloat(key string, value float64) *redis.FloatCmd {
	panic("implement me")
}

func (c *Client) MGet(keys ...string) *redis.SliceCmd {
	panic("implement me")
}

func (c *Client) MSet(pairs ...interface{}) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) MSetNX(pairs ...interface{}) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.client.Set(c.namespace+key, value, expiration)
}

func (c *Client) SetBit(key string, offset int64, value int) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) SetXX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) SetRange(key string, offset int64, value string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) StrLen(key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) HDel(key string, fields ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) HExists(key, field string) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) HGet(key, field string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) HGetAll(key string) *redis.StringStringMapCmd {
	panic("implement me")
}

func (c *Client) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	panic("implement me")
}

func (c *Client) HKeys(key string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) HLen(key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) HMGet(key string, fields ...string) *redis.SliceCmd {
	panic("implement me")
}

func (c *Client) HMSet(key string, fields map[string]interface{}) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) HSet(key, field string, value interface{}) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) HVals(key string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) LIndex(key string, index int64) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) LLen(key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) LPop(key string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) LPush(key string, values ...interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) LPushX(key string, value interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) LTrim(key string, start, stop int64) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) RPop(key string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) RPopLPush(source, destination string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) RPush(key string, values ...interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) RPushX(key string, value interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) SAdd(key string, members ...interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) SCard(key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) SDiff(keys ...string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) SDiffStore(destination string, keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) SInter(keys ...string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) SInterStore(destination string, keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) SIsMember(key string, member interface{}) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) SMembers(key string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) SMembersMap(key string) *redis.StringStructMapCmd {
	panic("implement me")
}

func (c *Client) SMove(source, destination string, member interface{}) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) SPop(key string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) SPopN(key string, count int64) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) SRandMember(key string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) SRem(key string, members ...interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) SUnion(keys ...string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) SUnionStore(destination string, keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) XAdd(a *redis.XAddArgs) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) XDel(stream string, ids ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) XLen(stream string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) XRange(stream, start, stop string) *redis.XMessageSliceCmd {
	panic("implement me")
}

func (c *Client) XRangeN(stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	panic("implement me")
}

func (c *Client) XRevRange(stream string, start, stop string) *redis.XMessageSliceCmd {
	panic("implement me")
}

func (c *Client) XRevRangeN(stream string, start, stop string, count int64) *redis.XMessageSliceCmd {
	panic("implement me")
}

func (c *Client) XRead(a *redis.XReadArgs) *redis.XStreamSliceCmd {
	panic("implement me")
}

func (c *Client) XReadStreams(streams ...string) *redis.XStreamSliceCmd {
	panic("implement me")
}

func (c *Client) XGroupCreate(stream, group, start string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) XGroupCreateMkStream(stream, group, start string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) XGroupSetID(stream, group, start string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) XGroupDestroy(stream, group string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) XGroupDelConsumer(stream, group, consumer string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) XReadGroup(a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	panic("implement me")
}

func (c *Client) XAck(stream, group string, ids ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) XPending(stream, group string) *redis.XPendingCmd {
	panic("implement me")
}

func (c *Client) XPendingExt(a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	panic("implement me")
}

func (c *Client) XClaim(a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	panic("implement me")
}

func (c *Client) XClaimJustID(a *redis.XClaimArgs) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) XTrim(key string, maxLen int64) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) XTrimApprox(key string, maxLen int64) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) BZPopMax(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	panic("implement me")
}

func (c *Client) BZPopMin(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	panic("implement me")
}

func (c *Client) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZAddNX(key string, members ...redis.Z) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZAddXX(key string, members ...redis.Z) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZAddCh(key string, members ...redis.Z) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZAddNXCh(key string, members ...redis.Z) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZAddXXCh(key string, members ...redis.Z) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZIncr(key string, member redis.Z) *redis.FloatCmd {
	panic("implement me")
}

func (c *Client) ZIncrNX(key string, member redis.Z) *redis.FloatCmd {
	panic("implement me")
}

func (c *Client) ZIncrXX(key string, member redis.Z) *redis.FloatCmd {
	panic("implement me")
}

func (c *Client) ZCard(key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZCount(key, min, max string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZLexCount(key, min, max string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	panic("implement me")
}

func (c *Client) ZInterStore(destination string, store redis.ZStore, keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZPopMax(key string, count ...int64) *redis.ZSliceCmd {
	panic("implement me")
}

func (c *Client) ZPopMin(key string, count ...int64) *redis.ZSliceCmd {
	panic("implement me")
}

func (c *Client) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	panic("implement me")
}

func (c *Client) ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) ZRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	panic("implement me")
}

func (c *Client) ZRank(key, member string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZRem(key string, members ...interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	panic("implement me")
}

func (c *Client) ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) ZRevRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	panic("implement me")
}

func (c *Client) ZRevRank(key, member string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ZScore(key, member string) *redis.FloatCmd {
	panic("implement me")
}

func (c *Client) ZUnionStore(dest string, store redis.ZStore, keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) PFAdd(key string, els ...interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) PFCount(keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) PFMerge(dest string, keys ...string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) BgRewriteAOF() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) BgSave() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClientKill(ipPort string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClientKillByFilter(keys ...string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ClientList() *redis.StringCmd {
	panic("implement me")
}

func (c *Client) ClientPause(dur time.Duration) *redis.BoolCmd {
	panic("implement me")
}

func (c *Client) ClientID() *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ConfigGet(parameter string) *redis.SliceCmd {
	panic("implement me")
}

func (c *Client) ConfigResetStat() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ConfigSet(parameter, value string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ConfigRewrite() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) DBSize() *redis.IntCmd {
	panic("implement me")
}

func (c *Client) FlushAll() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) FlushAllAsync() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) FlushDB() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) FlushDBAsync() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) Info(section ...string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) LastSave() *redis.IntCmd {
	panic("implement me")
}

func (c *Client) Save() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) Shutdown() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ShutdownSave() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ShutdownNoSave() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) SlaveOf(host, port string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) Time() *redis.TimeCmd {
	panic("implement me")
}

func (c *Client) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	panic("implement me")
}

func (c *Client) EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	panic("implement me")
}

func (c *Client) ScriptExists(hashes ...string) *redis.BoolSliceCmd {
	panic("implement me")
}

func (c *Client) ScriptFlush() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ScriptKill() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ScriptLoad(script string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) DebugObject(key string) *redis.StringCmd {
	panic("implement me")
}

func (c *Client) Publish(channel string, message interface{}) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) PubSubChannels(pattern string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) PubSubNumSub(channels ...string) *redis.StringIntMapCmd {
	panic("implement me")
}

func (c *Client) PubSubNumPat() *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ClusterSlots() *redis.ClusterSlotsCmd {
	panic("implement me")
}

func (c *Client) ClusterNodes() *redis.StringCmd {
	panic("implement me")
}

func (c *Client) ClusterMeet(host, port string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterForget(nodeID string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterReplicate(nodeID string) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterResetSoft() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterResetHard() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterInfo() *redis.StringCmd {
	panic("implement me")
}

func (c *Client) ClusterKeySlot(key string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ClusterGetKeysInSlot(slot int, count int) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) ClusterCountFailureReports(nodeID string) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ClusterCountKeysInSlot(slot int) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) ClusterDelSlots(slots ...int) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterDelSlotsRange(min, max int) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterSaveConfig() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterSlaves(nodeID string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) ClusterFailover() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterAddSlots(slots ...int) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ClusterAddSlotsRange(min, max int) *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	panic("implement me")
}

func (c *Client) GeoPos(key string, members ...string) *redis.GeoPosCmd {
	panic("implement me")
}

func (c *Client) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *Client) GeoRadiusRO(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *Client) GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *Client) GeoRadiusByMemberRO(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	panic("implement me")
}

func (c *Client) GeoDist(key string, member1, member2, unit string) *redis.FloatCmd {
	panic("implement me")
}

func (c *Client) GeoHash(key string, members ...string) *redis.StringSliceCmd {
	panic("implement me")
}

func (c *Client) ReadOnly() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) ReadWrite() *redis.StatusCmd {
	panic("implement me")
}

func (c *Client) MemoryUsage(key string, samples ...int) *redis.IntCmd {
	panic("implement me")
}
